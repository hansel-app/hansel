<template>
  <div class="background-gradient">
    <BackSwipe />
    <Header title="Friend requests" />
    <div class="container">
      <div v-if="pendingFriendRequests.length > 0">
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
                @click="acceptFriendRequest(request.requester.userId)"
              >
                Accept
              </van-button>
              <van-button round @click="declineFriendRequest(request.requester.userId)" size="small">
                Decline
              </van-button>
            </div>
          </FriendCell>
        </van-cell-group>
      </div>
      <div v-else>
        <i>You have no pending requests</i> &#x1f4a9;
      </div>
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
import { Toast } from "vant";

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
    acceptFriendRequest(requesterId: number) {
      const acceptRequest = () => this.store.dispatch("acceptFriendRequest", requesterId);
      acceptRequest()
        .then(() => {
          Toast.success("Friend request accepted!");
        })
        .catch((err) => Toast.fail(`Failed to accept request: ${err.message}`));
    },
    declineFriendRequest(requesterId: number) {
      const declineRequest = () => this.store.dispatch("declineFriendRequest", requesterId);
      declineRequest()
        .then(() => {
          Toast.success("Friend request declined!");
        })
        .catch((err) => Toast.fail(`Failed to decline request: ${err.message}`));
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
