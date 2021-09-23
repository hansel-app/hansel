<template>
  <div>
    <GemHeader title="Drop a gem" :color="color" />
    <div class="container">
      <p class="sub-header">Attach a message!</p>
      <Field
        v-model="message"
        class="message-input"
        placeholder="Enter your message"
        type="textarea"
        rows="5"
        @input="this.$emit('set-message-event', $event.target.value)"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useStore } from "vuex";
import { Field } from "vant";
import GemHeader from "@/components/GemHeader.vue";
import { GemColor } from "@/interfaces";
import { mapState } from "vuex";

export default defineComponent({
  components: {
    Field,
    GemHeader,
  },
  setup() {
    const store = useStore();
    return {
      store,
    };
  },
  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      color: (state: any) => state.gems.dropGemFormState.color as GemColor,
    }),
    message(): string {
      return this.store.state.gems.dropGemFormState.message;
    },
  },
});
</script>

<style scoped lang="less">
.message-input {
  background-color: #e9e9e9;
  border-radius: 2em;
  box-shadow: @box-shadow-gradient;
}
.container {
  margin-top: -10px;
  padding-top: -20px;
}
</style>
