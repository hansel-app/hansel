<template>
  <div class="background-gradient">
    <BackSwipe />
    <Header title="Add friends" />
    <div class="container">
      <Searchbar v-model="searchQuery" placeholder="Search by username" />
      <div v-if="filteredUsers.length > 0">
        <CellGroup>
          <FriendCell
            v-for="user in filteredUsers"
            :key="user.userId"
            :friend="user"
            :shouldDisplayUsername="true"
          >
            <van-icon
              :name="require('@/assets/icons/user-plus.svg')"
              size="24"
              @click="() => addFriend(user)"
            />
          </FriendCell>
        </CellGroup>
      </div>
      <div v-else-if="searchQuery.length > 0" class="feedback-message">
        <i>No results found</i> &#x1f4a9;
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { CellGroup, Toast } from "vant";
import { mapState, useStore } from "vuex";
import { User } from "@/interfaces";
import BackSwipe from "@/components/BackSwipe.vue";
import Searchbar from "@/components/Searchbar.vue";
import FriendCell from "@/components/FriendCell.vue";
import Header from "@/components/Header.vue";
import { AddFriendRequestResponse, AddFriendRequestStatus } from "@/protobuf/user_pb";

export default defineComponent({
  components: {
    BackSwipe,
    CellGroup,
    FriendCell,
    Searchbar,
    Header,
  },
  data() {
    return {
      searchQuery: ref(""),
      users: [] as User[],
      store: useStore(),
    };
  },
  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      friends: (state: any) => state.user.friends as User[],
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      pendingFriends: (state: any) => state.user.pendingFriends as User[],
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      selfUser: (state: any) => state.user.self as User,
    }),
    filteredUsers(): User[] {
      if (!this.searchQuery) {
        // Don't display any users when search bar is empty.
        return [];
      }
      this.handleSearch();
      return this.users as User[];
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    handleSearch() {
      const filterFriends = (user: User) => !this.friends.map(x => x.userId).includes(user.userId);
      const filterPendingFriends = (user: User) => !this.pendingFriends.map(x => x.userId).includes(user.userId);
      const filterSelf = (user: User) => this.selfUser.userId != user.userId;
      this.store
        .dispatch("searchByUsername", this.searchQuery)
        .then((users: User[]) => {
          this.users = users.filter(filterFriends).filter(filterPendingFriends).filter(filterSelf);
        });
    },
    addFriend(user: User) {
      this.store
        .dispatch("sendFriendRequest", user.userId)
        .then((resp: AddFriendRequestResponse) => {
          const status = resp.getStatus();
          if (status == AddFriendRequestStatus.SENT_FRIEND_REQUEST) {
            Toast.success(`Successfully sent friend request to @${user.username}`);
          } else if (status == AddFriendRequestStatus.ADDED_AS_FRIEND) {
            Toast.success(`Successfully added @${user.username} as friend`);
          } else {
            console.error(`Invalid status: ${status}`);
          }
          this.handleSearch();
        })
        .catch((err) => {
          Toast.fail(
            `Failed to send friend request to @${user.username}. Try again later!`
          );
          console.error(err);
        });
    },
  },
});
</script>
