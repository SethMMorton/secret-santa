<template>
  <q-item>
    <q-item-section>
      <q-expansion-item expand-separator :label="gifter">
        <q-option-group
          v-model="store.gifterToAllowedRecipients[gifter]"
          :options="options"
          type="checkbox"
        ></q-option-group>
      </q-expansion-item>
    </q-item-section>
    <q-item-section>
      <q-btn
        class="gt-xs"
        flat
        dense
        icon="delete"
        @click="removeParticipant"
      ></q-btn>
    </q-item-section>
  </q-item>
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

  methods: {
    removeParticipant() {
      this.store.removeParticipant(this.gifter);
    },
  },
});
</script>
