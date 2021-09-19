<template>
  <NavBar left-arrow @click-left="goBack" />
  <AvatarHeader
    :avatarUrl="placeholderAvatarUrl"
    :displayName="mockSelfUser.displayName"
    :username="mockSelfUser.username"
  />
  <h3 class="sub-header">Friends</h3>
  <CellGroup>
    <FriendCell v-for="user in mockFriends" :key="user.id" :friend="user" />
  </CellGroup>
  <Button round type="primary" @click="goToAddFriends">
    <Icon name="plus" />
  </Button>
  <!-- TODO: find out where to put friend requests button -->
  <Button type="primary" @click="goToFriendRequests">friend requests</Button>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Button, Icon, NavBar } from "vant";
import { mockSelfUser, mockFriends } from "@/interfaces/mockData";
import { ADD_FRIENDS_ROUTE, FRIEND_REQUESTS_ROUTE } from "@/constants";
import AvatarHeader from "./AvatarHeader.vue";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
  data() {
    return {
      // TODO: replace with user's own data
      placeholderAvatarUrl:
        "https://media.istockphoto.com/vectors/happy-young-woman-watching-into-rounded-frame-isolated-on-white-3d-vector-id1296058958?b=1&k=20&m=1296058958&s=170667a&w=0&h=6m2FU2hKv6emHjNtdNSqBJR1uMq64smptqwDAZNo6bg=",
      mockSelfUser,
      mockFriends,
    };
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    goToAddFriends() {
      this.$router.push(ADD_FRIENDS_ROUTE);
    },
    goToFriendRequests() {
      this.$router.push(FRIEND_REQUESTS_ROUTE);
    },
  },
  components: {
    AvatarHeader,
    Button,
    FriendCell,
    Icon,
    NavBar,
  },
});
</script>

<style scoped lang="less">
.sub-header {
  text-align: left;
}
</style>
