<template>
  <div class="background-gradient">
    <BackSwipe />
    <Header title="Add friends" />
    <div class="container">
      <Searchbar
        placeholder="Search by username"
        v-model="searchQuery"
        @input="handleSearch"
      />
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
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Search, Toast } from "vant";
import { useStore, mapState } from "vuex";
import { User } from "@/interfaces";
import BackSwipe from "@/components/BackSwipe.vue";
import Searchbar from "@/components/Searchbar.vue";
import FriendCell from "@/components/FriendCell.vue";
import Header from "@/components/Header.vue";

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
      searchQuery: "",
      users: [] as User[],
      store: useStore(),
    };
  },
  computed: {
    ...mapState({
      friends: (state: any) => state.user.friends as User[],
      selfUser: (state: any) => state.user.self as User,
    }),
    filteredUsers(): User[] {
      if (!this.searchQuery) {
        // Don't display any users when search bar is empty.
        return [];
      }
      return this.users as User[];
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    handleSearch() {
      const filterFriends = (user: User) => !this.friends.includes(user);
      const filterSelf = (user: User) => this.selfUser.userId != user.userId;
      this.store
        .dispatch("searchByUsername", this.searchQuery)
        .then((users: User[]) => {
          this.users = users.filter(filterFriends).filter(filterSelf);
        });
    },
    addFriend(user: User) {
      this.store
        .dispatch("sendFriendRequest", user.userId)
        .then(() =>
          Toast.success(`Successfully sent friend request to @${user.username}`)
        )
        .catch((err) => {
          Toast.success(
            `Failed to send friend request to @${user.username}. Try again later!`
          );
          console.error(err);
        });
    },
  },
});
</script>
