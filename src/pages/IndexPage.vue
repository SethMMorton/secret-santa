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
      <q-list>
        <ParticipantProperties
          v-for="k in store.allParticipants()"
          :key="k"
          :gifter="k"
        >
        </ParticipantProperties>
      </q-list>
      <div>
        <q-btn label="Submit" type="submit" color="primary"></q-btn>
        <q-btn label="Reset" type="reset" color="primary" flat></q-btn>
      </div>
    </q-form>
    <q-table
      v-if="Object.keys(solution).length > 0"
      title="Secret Santa Solution"
      :rows="solutionRows"
      :columns="solutionColumns"
      row-key="from"
      separator="vertical"
    ></q-table>
  </q-page>
</template>

<script>
import { defineComponent } from "vue";

import { store } from "../store";

import ParticipantProperties from "components/ParticipantProperties.vue";

function shuffleArray(array) {
  for (let i = array.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [array[i], array[j]] = [array[j], array[i]];
  }
}

function factorial(n) {
  let ans = 1;
  for (let i = 2; i <= n; i++) {
    ans = ans * i;
  }
  return ans;
}

// Determine if any gifters are giving to someone who
// is also giving back to them.
function circularPairing(giftersToRecievers) {
  for (const firstPerson in giftersToRecievers) {
    const secondPerson = giftersToRecievers[firstPerson];
    const thirdPerson = giftersToRecievers[secondPerson];
    if (firstPerson === thirdPerson) {
      return true;
    }
  }
  return false;
}

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
      solutionColumns: [
        {
          name: "gifter",
          required: true,
          label: "From",
          align: "left",
          field: "from",
          sortable: true,
        },
        {
          name: "recipient",
          required: true,
          label: "To",
          align: "left",
          field: "to",
          sortable: true,
        },
      ],
    };
  },

  computed: {
    solutionRows() {
      console.log(
        Object.entries(this.solution).map(([gifter, receiver]) => ({
          from: gifter,
          to: receiver,
        }))
      );
      return Object.entries(this.solution).map(([gifter, receiver]) => ({
        from: gifter,
        to: receiver,
      }));
    },
  },

  methods: {
    addParticipant() {
      if (this.participant !== "") {
        this.store.addParticipant(this.participant);
        this.participant = "";
      }
    },

    resetParticipants() {
      this.participant = "";
      this.solution = {};
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
