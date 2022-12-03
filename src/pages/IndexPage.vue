<template>
  <q-page class="flex flex-center">
    <q-form @submit="calculate" @reset="resetParticipants">
      <q-input
        v-model="participant"
        label="Add a participant"
        hint="Press enter to add"
        clearable
        outlined
        @keydown.enter.prevent="addParticipant"
      >
      </q-input>
      <div>
        <q-btn label="Submit" type="submit" color="primary"></q-btn>
        <q-btn label="Reset" type="reset" color="primary" flat></q-btn>
      </div>
    </q-form>
    <q-list>
      <ParticipantProperties
        v-for="k in store.allParticipants()"
        :key="k"
        :gifter="k"
      >
      </ParticipantProperties>
    </q-list>
  </q-page>
  <q-input :model-value="results" readonly autogrow></q-input>
</template>

<script>
import { defineComponent } from "vue";

import { store } from "../store";
import { factorial, circularPairing, shuffleArray } from "../util";

import ParticipantProperties from "components/ParticipantProperties.vue";

export default defineComponent({
  name: "IndexPage",

  components: {
    ParticipantProperties,
  },

  data() {
    return {
      participant: "",
      solution: {},
      store,
    };
  },

  computed: {
    // Format the recipients list nicely
    results() {
      const gifters = Object.keys(this.solution);
      if (gifters.length === 0) {
        return "";
      }
      const lengths = gifters.map((value) => value.length);
      const maxLen = Math.max(...lengths);

      let resultArray = [];
      for (const [gifter, recipient] of Object.entries(this.solution)) {
        resultArray.push(gifter.padEnd(maxLen));
        resultArray.push(" ==> ");
        resultArray.push(recipient);
        resultArray.push("\n");
      }
      console.log(resultArray.join(""));
      return resultArray.join("");
    },
  },

  methods: {
    addParticipant() {
      if (this.participant !== "") {
        this.store.addParticipant(this.participant);
      }
    },
    resetParticipants() {
      this.participant = "";
      this.store.resetParticipants();
    },
    calculate() {
      // Initiallize with the given list
      const gifters = this.store.allParticipants();
      let recipients = gifters.slice(0); // clone array

      // Iterate until we get a solution that works
      // up to the total number of possible combinations.
      // This is a "brute force" solution - really, not smart
      let seen = new Set();
      const combinations = factorial(gifters.length);
      while (seen.size < combinations) {
        // Shuffle the recipeients.
        shuffleArray(recipients);

        // Keep track of if we have seen this combination
        // and skip if we have already.
        const key = recipients.join(" ");
        if (seen.has(key)) {
          continue;
        }
        seen.add(key);

        // Zip the gifters and recipients arrays into a mapping.
        const tentativeSolution = Object.fromEntries(
          gifters.map((_, i) => [gifters[i], recipients[i]])
        );

        // If any pairing is self or invalid, reject this selection.
        let disallowedSolution = false;
        for (const [gifter, recipient] of Object.entries(tentativeSolution)) {
          // Check if this solution is disallowed by the rules
          disallowedSolution = this.store.isDisallowed(gifter, recipient);
          if (disallowedSolution) {
            break;
          }
        }

        // If solution is disallowed by the rules, or contains a circular pairing,
        // we discard it.
        if (disallowedSolution || circularPairing(tentativeSolution)) {
          continue;
        }

        // If we got here, then this is a valid solution!
        this.solution = tentativeSolution;
        return;
      }

      // Invalid solution, return empty array
      this.solution = {};
    },
  },
});
</script>
