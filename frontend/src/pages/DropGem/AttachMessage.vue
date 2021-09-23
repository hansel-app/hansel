<template>
  <div>
    <GemHeader title="Drop a gem" :color="color" />
    <div class="container">
      <p class="sub-header">Attach a message!</p>
      <Field
        class="message-input"
        placeholder="Enter your message"
        type="textarea"
        rows="5"
        v-model="message"
        @input="this.$emit('set-message-event', $event.target.value)"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Field } from "vant";
import GemHeader from "@/components/GemHeader.vue";
import { GemColor } from "@/interfaces";
import { mapState } from "vuex";

export default defineComponent({
  components: {
    Field,
    GemHeader,
  },
  data() {
    return {
      message: "",
    };
  },
  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      color: (state: any) => state.gems.dropGemFormState.color as GemColor,
    }),
  },
});
</script>

<style scoped lang="less">
.message-input {
  background-color: #e9e9e9;
}
</style>
