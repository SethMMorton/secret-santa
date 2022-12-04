<template>
  <q-page class="row">
    <q-card class="col-4 q-ma-xl">
      "Secret Santa" is a gift-giving game wherin each participant (the giver)
      is assigned the name of another participant (the reciever), and the giver
      purchases a gift for the reciever. Each person gets a gift! For large
      families or work settings, it is a good way to share gifts without making
      people feel as though they must spend a lot of time and money purchasing
      gifts for everyone.
      <br />
      <br />
      Traditionally, at some point ahead of the gathering, names are placed into
      a hat and each participant draws for whom they will be purchasing.
      However, there are some times when getting everyone together is not
      feasable (for example, in a global pandemic), so this program serves as a
      virtual "hat". To be fair - this program does not generate secret results.
      At least the person executing the program will see all the assignments...
      <br />
      <br />
      However, in practice, it is often desirable that people that are very
      close (e.g. spouses, significant others, etc.) not get each other in the
      Secret Santa because they were already buying each other gifts. To solve
      this problem you can assign a list of "invalid" recipients for each
      partipant. The `secret-santa` program will not allow these people to be
      recipients for that participant.
    </q-card>
    <q-card class="col-3 q-ma-xl">
      <ParticipantInput
        @solution="(solution) => (this.solution = solution)"
      ></ParticipantInput>
    </q-card>
    <q-table
      v-if="Object.keys(solution).length > 0"
      class="col q-ma-xl"
      title="Secret Santa Solution"
      :rows="solutionRows"
      :columns="solutionColumns"
      row-key="from"
      separator="vertical"
      :pagination="{ rowsPerPage: Object.keys(solution).length }"
      hide-pagination
    ></q-table>
  </q-page>
</template>

<script>
import { defineComponent } from "vue";

import ParticipantInput from "components/ParticipantInput.vue";

export default defineComponent({
  name: "IndexPage",

  components: { ParticipantInput },

  data() {
    return {
      solution: {},
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
      return Object.entries(this.solution).map(([gifter, receiver]) => ({
        from: gifter,
        to: receiver,
      }));
    },
  },
});
</script>
