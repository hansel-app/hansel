<template>
  <div>
    <div class="backswipe-area" v-touch:swipe.right="prevStage"></div>
    <div v-if="currentStage > 0">
      <img id="icon-left" @click="prevStage" :src="arrowLeft" />
    </div>
    <router-view
      :name="currentStageName"
      @next-stage="nextStage"
      @set-receiver-event="setReceiverId"
      @set-message-event="setMessage"
      @set-gem-color-event="setGemColor"
    ></router-view>

    <van-button
      class="form-navigation-button"
      round
      type="primary"
      v-if="shouldShowNextButton"
      @click="nextStage"
    >
      Next
    </van-button>
    <van-button
      class="form-navigation-button"
      round
      type="primary"
      v-if="currentStage === numStages - 1"
      v-on:click="dropMyGem"
    >
      Drop Gem
    </van-button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, Ref } from "vue";
import { Toast } from "vant";
import { useStore } from "vuex";
import { GemColor } from "@/interfaces";
import { HOME_ROUTE } from "@/constants";
import arrowLeft from "@/assets/icons/arrow-left.svg";

enum DropGemStage {
  Friend,
  Media,
  Message,
}

const NO_PHOTO_ERROR = "You must attach a photo";
const NO_MESSAGE_ERROR = "You must write something";

export default defineComponent({
  setup() {
    const store = useStore();
    store.commit("clearDropGemFormState");

    const currentStage: Ref<DropGemStage> = ref(DropGemStage.Friend);

    return {
      currentStage,
      store,
      arrowLeft,
    };
  },
  computed: {
    numStages(): number {
      return Object.keys(DropGemStage).length / 2;
    },
    currentStageName(): string {
      return DropGemStage[this.currentStage];
    },
    shouldShowNextButton(): boolean {
      return [DropGemStage.Media].includes(this.currentStage);
    },
    isMediaAttached(): boolean {
      return this.store.state.gems.dropGemFormState.attachment !== undefined;
    },
    isMessageFilled(): boolean {
      return this.store.state.gems.dropGemFormState.message;
    },
  },
  methods: {
    nextStage() {
      if (this.currentStage === DropGemStage.Media && !this.isMediaAttached) {
        Toast.fail(NO_PHOTO_ERROR);
        return;
      }
      if (this.currentStage < this.numStages - 1) {
        this.currentStage += 1;
      }
    },
    prevStage() {
      if (this.currentStage > 0) {
        this.currentStage -= 1;
      }
    },
    setReceiverId(value: number): void {
      this.store.commit("updateDropGemFormState", { receiverId: value });
    },
    setMessage(value: string): void {
      this.store.commit("updateDropGemFormState", { message: value });
    },
    setGemColor(value: GemColor): void {
      this.store.commit("updateDropGemFormState", { color: value });
    },
    dropMyGem() {
      // Should not happen, but just in case.
      if (!this.isMediaAttached) {
        Toast.fail(NO_PHOTO_ERROR);
        return;
      }
      if (!this.isMessageFilled) {
        Toast.fail(NO_MESSAGE_ERROR);
        return;
      }
      Toast.loading({ message: "Loading...", duration: 0, forbidClick: true });
      const dropGem = () => this.store.dispatch("dropGem");
      dropGem()
        .then(() => {
          Toast.success("Dropped a gem!");
          this.$router.replace(HOME_ROUTE);
        })
        .catch((err) => Toast.fail(`Failed to drop gem: ${err.message}`));
    },
  },
});
</script>

<style scoped>
.form-navigation-button {
  width: 80vw;
}
#icon-left {
  position: absolute;
  left: 1em;
  top: 2.8em;
  z-index: 1;
}
</style>
