<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-toolbar-title>Secret Santa Solver</q-toolbar-title>
        <div>Version {{ version }}</div>
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
