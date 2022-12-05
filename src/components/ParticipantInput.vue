<template>
  <q-form @submit="calculate" @reset="resetParticipants">
    <q-input
      class="q-pa-sm"
      v-model="participant"
      label="Add a participant"
      hint="Press enter to add"
      clearable
      outlined
      @keydown.enter.prevent="addParticipant"
    >
    </q-input>
    <q-space class="q-pa-md"></q-space>
    <div class="q-gutter-sm q-pa-sm flex justify-around">
      <q-btn
        label="Solve The Secret Santa Problem"
        type="submit"
        color="primary"
      ></q-btn>
      <q-btn label="Start Over" type="reset" color="secondary"></q-btn>
      <q-checkbox
        :model-value="expanded"
        indeterminate-value="some"
        toggle-order="ft"
        @update:modelValue="setExpanded"
      >
        <q-tooltip class="text-body2">
          Expand or collapse all potential reciever lists for each gifter
        </q-tooltip>
      </q-checkbox>
    </div>
    <q-list class="q-pa-sm">
      <ParticipantProperties
        v-for="k in store.allParticipants()"
        v-model="expandedParticipant[k]"
        :key="k"
        :gifter="k"
        @participant-removed="$emit('solution', {})"
        @update:modelValue="updateExpanded"
      >
      </ParticipantProperties>
    </q-list>
  </q-form>
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
  name: "ParticipantInput",

  components: { ParticipantProperties },

  emits: ["solution", "cannotSolve"],

  data() {
    return {
      participant: "",
      store,
      expandedParticipant: {},
    };
  },

  computed: {
    expanded() {
      const allClosed = Object.entries(this.expandedParticipant).every(
        ([k, v]) => !v
      );
      const allOpened = Object.entries(this.expandedParticipant).every(
        ([k, v]) => v
      );
      if (allClosed) {
        return false;
      } else if (allOpened) {
        return true;
      } else {
        return "some";
      }
    },
  },

  methods: {
    addParticipant() {
      if (this.participant !== "") {
        this.store.addParticipant(this.participant);
        this.expandedParticipant[this.participant] = false;
        this.participant = "";
        this.$emit("solution", {});
        this.updateExpanded();
      }
    },

    resetParticipants() {
      this.participant = "";
      this.store.resetParticipants();
      this.$emit("solution", {});
      delete this.expandedParticipant[this.participant];
      this.updateExpanded();
    },

    setExpanded(value) {
      Object.keys(this.expandedParticipant).forEach(
        (k) => (this.expandedParticipant[k] = value)
      );
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
        this.$emit("solution", tentativeSolution);
        return;
      }

      // Invalid solution, return empty solution
      this.$emit("solution", {});
      this.$emit("cannotSolve");
    },
  },
});
</script>
