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
        <van-button round type="primary" @click="acceptFriendRequest"
          >Accept</van-button
        >
        <van-button round @click="declineFriendRequest">Decline</van-button>
      </div>
    </FriendCell>
  </van-cell-group>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mockFriends } from "@/interfaces/mockData";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
  components: {
    FriendCell,
  },
  data() {
    return {
      mockFriends,
    };
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    acceptFriendRequest() {
      console.log("accept friend request");
    },
    declineFriendRequest() {
      console.log("decline friend request");
    },
  },
});
</script>
