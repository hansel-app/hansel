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
import { defineComponent } from "vue";
import { mockFriends } from "@/interfaces/mockData";
import FriendCell from "@/components/FriendCell.vue";
import Header from "@/components/Header.vue";

export default defineComponent({
  components: {
    FriendCell,
    Header,
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
<style scoped>
#left-button {
  margin-right: 1em;
}
</style>