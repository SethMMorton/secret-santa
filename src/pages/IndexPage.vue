<template>
  <q-page class="row">
    <q-card class="col-4 q-ma-xl">
      <q-card-section>
        <div class="text-h6">What is Secret Santa?</div>
      </q-card-section>
      <q-card-section>
        "Secret Santa" is a gift-giving game wherin each participant (the giver)
        is assigned the name of another participant (the reciever), and the
        giver purchases a gift for the reciever. Each person gets a gift! For
        large families or work settings, it is a good way to share gifts without
        making people feel as though they must spend a lot of time and money
        purchasing gifts for everyone.
      </q-card-section>
      <q-card-section>
        Traditionally, at some point ahead of the gathering, names are placed
        into a hat and each participant draws for whom they will be purchasing.
        However, there are some times when getting everyone together is not
        feasable (for example, in a global pandemic), so this program serves as
        a virtual "hat".
      </q-card-section>
      <q-card-section>
        To be fair - this program does not generate secret results... at a
        minimum the person executing the program will see all the assignments.
      </q-card-section>
      <q-separator></q-separator>
      <q-card-section>
        <div class="text-h6">Instructions</div>
        <ol>
          <li>
            Add the names of all the participants in the Secret Santa by typing
            thier name into the "Add a Participant" field and pressing "Enter".
          </li>
          <ul>
            <li>
              If at any time you need to remove a name, press the red "Trash"
              icon to the right of thier name to remove them.
            </li>
          </ul>
          <li>
            (<strong>Optional</strong>) You can specify which recipients are
            valid for each gifter by pressing the down arrow to the right of
            each name and checking or unchecking names - checked names are
            possible recipients.
            <ul>
              <li>
                This can be used to prevent people that are very close (e.g.
                spouses, significant others, etc.) from getting each other in
                the Secret Santa, which could be a problem because they were
                already buying each other gifts.
              </li>
            </ul>
          </li>
          <li>
            Press the green "Solve The Secret Santa Problem" to see the solution
            for your Secret Santa criteria.
          </li>
          <ul>
            <li>
              The solution will be presented in a table on the right side of the
              screen.
            </li>
          </ul>
        </ol>
        You can start over at any time by pressing the red "Start Over" button.
      </q-card-section>
    </q-card>
    <q-card class="col-3 q-ma-xl">
      <ParticipantInput
        @solution="(solution) => (this.solution = solution)"
        @cannot-solve="solverError = true"
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
    <q-dialog v-model="solverError">
      <q-card class="bg-secondary text-white">
        <q-card-section>
          <div class="text-h6">Cannot Solve Your Secret Santa!</div>
        </q-card-section>
        <q-card-section class="q-pt-none">
          Sorry! The criteria you have provided is not solvable! Some
          troubleshooting tips:
          <ul>
            <li>There must be three or more participants.</li>
            <li>
              Ensure that each gifter has at <em>least</em> one valid recipient.
            </li>
            <li>
              The more valid recipients each gifter has available, the more
              solutions will present themselves.
            </li>
          </ul>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn label="Got It" color="primary" v-close-popup></q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>
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
      solverError: false,
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
