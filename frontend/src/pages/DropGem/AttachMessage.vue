<template>
  <GemHeader title="Drop a gem" />
  <h4>Attach a message!</h4>
  <Field
    class="message-input"
    placeholder="Enter your message"
    type="textarea"
    rows="5"
  />
  <Button v-on:click="dropGem">Drop gem</Button>
</template>

<script>
import { defineComponent } from "vue";
import { Button, Field } from "vant";
import GemHeader from "@/components/GemHeader.vue";

export default defineComponent({
  components: {
    Button,
    Field,
    GemHeader,
  },
  methods: {
    dropGem: function() {
      console.log("Dropping gem");
      this.$router.replace({ name: "home" });
    },
  },
});
</script>

<style scoped>
.message-input {
  background-color: #e9e9e9;
}
</style>
