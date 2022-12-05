<template>
  <q-item class="row">
    <q-item-section class="col">
      <q-expansion-item v-model="expanded" expand-separator :label="gifter">
        <q-tooltip class="text-body2">
          Disallow this receiver for this gifter
        </q-tooltip>
        <q-option-group
          v-model="store.gifterToAllowedRecipients[gifter]"
          :options="options"
          type="checkbox"
        ></q-option-group>
      </q-expansion-item>
    </q-item-section>
    <q-item-section class="col-1">
      <q-btn
        class="gt-xs"
        flat
        dense
        icon="delete"
        color="secondary"
        @click="removeParticipant"
      >
        <q-tooltip class="text-body2">Remove this participant</q-tooltip>
      </q-btn>
    </q-item-section>
  </q-item>
</template>

<script>
import { defineComponent } from "vue";

import { store } from "../store";

export default defineComponent({
  name: "ParticipantProperties",

  emits: ["participantRemoved", "update:modelValue"],

  props: {
    gifter: {
      type: String,
      required: true,
    },
    modelValue: {
      type: Boolean,
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
    expanded: {
      get() {
        return this.modelValue;
      },
      set(value) {
        this.$emit("update:modelValue", value);
      },
    },
  },

  methods: {
    removeParticipant() {
      this.store.removeParticipant(this.gifter);
      this.$emit("participantRemoved");
    },
  },
});
</script>
