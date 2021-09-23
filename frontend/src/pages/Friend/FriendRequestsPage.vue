<template>
  <div class="background-gradient">
    <BackSwipe />
    <Header title="Friend requests" />
    <div class="container">
      <van-cell-group>
        <FriendCell
          v-for="request in pendingFriendRequests"
          :key="request.requester.userId"
          :friend="request.requester"
          :shouldDisplayUsername="true"
        >
          <div class="align-right">
            <van-button
              id="left-button"
              round
              type="primary"
              size="small"
              @click="acceptFriendRequest"
            >
              Accept
            </van-button>
            <van-button round @click="declineFriendRequest" size="small">
              Decline
            </van-button>
          </div>
        </FriendCell>
      </van-cell-group>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import BackSwipe from "@/components/BackSwipe.vue";
import FriendCell from "@/components/FriendCell.vue";
import { useStore } from "vuex";
import Header from "@/components/Header.vue";
import { PendingFriendRequest } from "@/interfaces";

export default defineComponent({
  setup() {
    const store = useStore();
    const fetchFriendRequests = () => store.dispatch("getFriendRequests");

    onMounted(() => fetchFriendRequests);
  },
  components: {
    BackSwipe,
    FriendCell,
    Header,
  },
  data() {
    return {
      store: useStore(),
    };
  },
  computed: {
    pendingFriendRequests(): PendingFriendRequest {
      return this.store.state.user.friendRequests;
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    acceptFriendRequest(senderId: number) {
      this.store.dispatch("acceptFriendRequest", senderId);
    },
    declineFriendRequest(senderId: number) {
      this.store.dispatch("declineFriendRequest", senderId);
    },
  },
});
</script>
<style scoped>
#left-button {
  margin-right: 0.5em;
}
.align-right {
  margin: auto 0 auto auto;
  position: absolute;
  right: 0;
}
</style>
