<template>
  <q-expansion-item expand-separator :label="gifter">
    <q-option-group
      v-model="store.gifterToAllowedRecipients[gifter]"
      :options="options"
      type="checkbox"
    ></q-option-group>
  </q-expansion-item>
</template>

<script>
import { defineComponent } from "vue";

import { store } from "../store";

export default defineComponent({
  name: "ParticipantProperties",

  props: {
    gifter: {
      type: String,
      required: true,
    },
  },

  data() {
    return {
      store,
    };
  },

  computed: {
    // All participants, formatted the way the option group wants them
    options() {
      return this.store
        .allParticipants()
        .filter((recipent) => recipent !== this.gifter)
        .map((v) => ({ label: v, value: v }));
    },
  },
});
</script>
