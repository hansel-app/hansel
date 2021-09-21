<template>
  <van-nav-bar left-arrow @click-left="goBack" />
  <h2>friend requests</h2>
  <van-cell-group>
    <FriendCell
      v-for="user in mockFriends"
      :key="user.id"
      :friend="user"
      :shouldDisplayUsername="true"
    >
      <div>
        <van-button round type="primary" @click="acceptFriendRequest(user.id)">
          Accept
        </van-button>
        <van-button round @click="declineFriendRequest(user.id)"
          >Decline</van-button
        >
      </div>
    </FriendCell>
  </van-cell-group>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { mockFriends } from "@/interfaces/mockData";
import { PersonInfo } from "@/protobuf/user_pb";
import FriendCell from "@/components/FriendCell.vue";
import { useStore } from "vuex";

export default defineComponent({
  setup() {
    const store = useStore();
    const fetchFriendRequests = () => store.dispatch("getFriendRequests");

    onMounted(() => fetchFriendRequests);
  },
  components: {
    FriendCell,
  },
  data() {
    return {
      mockFriends,
      store: useStore(),
    };
  },
  computed: {
    // TODO: replace mockFriends with this.
    // Not sure how to update the seed haha so I'll leave it.
    pendingFriendRequests(): PersonInfo[] {
      return this.store.state.user.friendRequests;
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    acceptFriendRequest(sender_id: number) {
      console.log("accept friend request from " + sender_id);
      this.store.dispatch("acceptFriendRequest", sender_id);
    },
    declineFriendRequest(sender_id: number) {
      console.log("decline friend request from " + sender_id);
      this.store.dispatch("declineFriendRequest", sender_id);
    },
  },
});
</script>
