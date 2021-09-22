<template>
  <div class="background-gradient">
    <Header title="Friend requests"/>
    <div class="container">
      <van-cell-group>
        <FriendCell
          v-for="user in mockFriends"
          :key="user.id"
          :friend="user"
          :shouldDisplayUsername="true"
        >
          <div>
            <van-button id="left-button"
              round 
              type="primary" 
              size="small"
              @click="acceptFriendRequest"
            >Accept</van-button>
            <van-button 
              round @click="declineFriendRequest"
              size="small"
            >Decline</van-button>
          </div>
        </FriendCell>
      </van-cell-group>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { mockFriends } from "@/interfaces/mockData";
import { PersonInfo } from "@/protobuf/user_pb";
import FriendCell from "@/components/FriendCell.vue";
import { useStore } from "vuex";
import Header from "@/components/Header.vue";

export default defineComponent({
  setup() {
    const store = useStore();
    const fetchFriendRequests = () => store.dispatch("getFriendRequests");

    onMounted(() => fetchFriendRequests);
  },
  components: {
    FriendCell,
    Header,
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
  margin-right: 1em;
}
</style>