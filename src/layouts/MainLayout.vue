<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-toolbar-title>Secret Santa Solver</q-toolbar-title>
        <div>
          <strong>Version {{ version }}</strong>
        </div>
        <q-btn
          class="q-pa-sm"
          size="md"
          color="secondary"
          icon="report_problem"
          href="https://github.com/SethMMorton/secret-santa/issues"
          flat
        >
          <q-tooltip>Report a problem</q-tooltip>
        </q-btn>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { defineComponent } from "vue";
import { version } from "../../package.json";

export default defineComponent({
  name: "MainLayout",

  data() {
    return { version: version };
  },
});
</script>
