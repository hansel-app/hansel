<template>
  <GemHeader title="Drop a gem" />
  <div class="container">
    <p class="sub-header">Attach a message!</p>
    <Field
      class="message-input"
      placeholder="Enter your message"
      type="textarea"
      rows="5"
    />
  </div>
  <Button v-on:click="dropGem">Drop gem</Button>
</template>

<script>
import { defineComponent } from "vue";
import { Button, Field } from "vant";
import GemHeader from "@/components/GemHeader.vue";
import { HOME_ROUTE } from "@/constants";

export default defineComponent({
  components: {
    Button,
    Field,
    GemHeader,
  },
  methods: {
    dropGem: function() {
      console.log("Dropping gem");
      this.$router.replace(HOME_ROUTE);
    },
  },
});
</script>

<style scoped lang="less">
.message-input {
  background-color: #e9e9e9;
}
</style>
