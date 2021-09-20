<template>
  <GemHeader title="Drop a gem" />
  <div class="container">
    <p class="sub-header">Attach a message!</p>
    <Field
      class="message-input"
      placeholder="Enter your message"
      type="textarea"
      rows="5"
      @input="this.$emit('SetMessageEvent', $event.target.value)"
    />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Field } from "vant";
import GemHeader from "@/components/GemHeader.vue";

export default defineComponent({
  components: {
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

<style scoped lang="less">
.message-input {
  background-color: #e9e9e9;
}
</style>
