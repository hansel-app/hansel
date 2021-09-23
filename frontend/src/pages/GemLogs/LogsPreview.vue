<template>
  <FriendCell
    :friend="friend"
    :isClickable="true"
    :labelOverride="previewMessage"
    @click="goToLogsPreview"
  />
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { mapState, useStore } from "vuex";
import { User, Gem, GemLogsWithFriend } from "@/interfaces";
import { GEM_LOGS_ROUTE } from "@/constants";
import FriendCell from "@/components/FriendCell.vue";

enum GemActivityPreviewMessage {
  SELF_DROPPED = "You dropped a gem!",
  SELF_PICKED_UP = "You picked up their gem!",
  PEER_DROPPED = "They dropped a gem!",
  PEER_PICKED_UP = "They picked up your gem!",
}

export default defineComponent({
  props: {
    friend: {
      type: Object as PropType<User>,
      required: true,
    },
    mostRecentGemActivity: {
      type: Object as PropType<Gem>,
      required: true,
    },
    gemLogsWithFriend: {
      type: Object as PropType<GemLogsWithFriend>,
      required: true,
    },
  },
  setup() {
    const store = useStore();

    return { store };
  },
  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      selfUser: (state: any) => state.user.self as User,
    }),
    previewMessage(): string {
      const isPickedUp = this.mostRecentGemActivity.receivedAt !== null;
      // TODO: replace hardcoded user value with this.store.user.self.id;
      const isSelfCreator =
        this.mostRecentGemActivity.createdBy.userId === this.selfUser.userId;

      if (isSelfCreator) {
        // I created and they picked up.
        if (isPickedUp) return GemActivityPreviewMessage.PEER_PICKED_UP;
        // I created and they have yet to pick up.
        else return GemActivityPreviewMessage.SELF_DROPPED;
      } else {
        // They created and I picked up.
        if (isPickedUp) return GemActivityPreviewMessage.SELF_PICKED_UP;
        // They created and I have yet to pick up.
        else return GemActivityPreviewMessage.PEER_DROPPED;
      }
    },
  },
  components: {
    FriendCell,
  },
  methods: {
    goToLogsPreview(): void {
      this.store.commit("setSelectedGemLog", this.gemLogsWithFriend);
      this.$router.push(`${GEM_LOGS_ROUTE}/${this.friend.userId}`);
    },
  },
});
</script>
