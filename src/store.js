import { reactive } from "vue";

// This object contains data that can be shared amongst components
export const store = reactive({
  // The object that maps a gifter to allowed recipients
  gifterToAllowedRecipients: {},

  // An array of all participants
  allParticipants() {
    return Object.keys(this.gifterToAllowedRecipients);
  },

  // Remove all participants
  resetParticipants() {
    for (const name of this.allParticipants()) {
      this.removeParticipant(name);
    }
  },

  // Add a participant - adds this participant to all other lists
  addParticipant(name) {
    if (!this.gifterToAllowedRecipients.hasOwnProperty(name)) {
      for (const others of this.allParticipants()) {
        this.allowRecipient(others, name);
      }
      this.gifterToAllowedRecipients[name] = this.allParticipants();
    }
  },

  // Remove a participant
  removeParticipant(name) {
    delete this.gifterToAllowedRecipients[name];
  },

  // The disallowed recipenents for a given gifter
  disallowedRecipients(gifter) {
    return this.allParticipants().filter((recipent) =>
      this.isDisallowed(gifter, recipent)
    );
  },

  // The allowed recipenents for a given gifter
  allowedRecipients(gifter) {
    return this.gifterToAllowedRecipients[gifter];
  },

  // Is this recipeient not allowed for this gifter?
  isDisallowed(gifter, recipent) {
    return !this.isAllowed(gifter, recipent);
  },

  // Is this recipeient allowed for this gifter?
  isAllowed(gifter, recipent) {
    return (
      this.gifterToAllowedRecipients[gifter].includes(recipent) &&
      recipent !== gifter
    );
  },

  // Set a recipent as disallowed for this gifter
  disallowRecipient(gifter, recipient) {
    if (gifter !== recipient) {
      for (const i in this.gifterToAllowedRecipients[gifter]) {
        if (this.gifterToAllowedRecipients[gifter][i] === recipient) {
          this.gifterToAllowedRecipients[gifter].splice(i, 1);
          break;
        }
      }
    }
  },

  // Set a recipent as allowed for this gifter
  allowRecipient(gifter, recipient) {
    const exists = this.gifterToAllowedRecipients[gifter].includes(recipient);
    if (gifter !== recipient && !exists) {
      this.gifterToAllowedRecipients[gifter].push(recipient);
    }
  },

  // Add a list of recipents as disallowed for this gifter
  disallowRecipients(gifter, recipientList) {
    for (const recipient of recipientList) {
      this.disallowRecipient(gifter, recipient);
    }
  },

  // Add a list of recipents as allowed for this gifter
  allowRecipients(gifter, recipientList) {
    for (const recipient of recipientList) {
      this.allowRecipient(gifter, recipient);
    }
  },

  // Make the recipients for a gifter match the recipient list
  setRecipients(gifter, recipientList) {
    this.gifterToAllowedRecipients[gifter].length = 0;
    this.allowRecipients(gifter, recipientList);
  },
});
