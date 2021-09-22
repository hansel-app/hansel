<template>
  <div class="background-gradient">
    <Header title="Add friends" />
    <div class="container">
      <Search
        placeholder="Search by username"
        v-model="searchQuery"
        @input="handleSearch"
      />
      <CellGroup>
        <FriendCell
          v-for="user in filteredUsers"
          :key="user.id"
          :friend="user"
          :shouldDisplayUsername="true"
        >
          <van-icon
            :name="require('@/assets/icons/user-plus.svg')"
            size="24"
            @click="(user) => addFriend(user)"
          />
        </FriendCell>
      </CellGroup>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Search } from "vant";
import { User } from "@/interfaces";
import { useStore, mapState } from "vuex";
import FriendCell from "@/components/FriendCell.vue";
import Header from "@/components/Header.vue";

export default defineComponent({
  components: {
    CellGroup,
    FriendCell,
    Header,
    Search,
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
      friends: (state: any) => state.user.friends,
      selfUser: (state: any) => state.user.self,
    }),
    filteredUsers(): User[] {
      if (!this.searchQuery) {
        // Don't display any users when search bar is empty.
        return [];
      }
      return this.users;
    },
  },
  methods: {
    goBack() {
      this.$router.back();
    },
    handleSearch() {
      const filterFriends = (user: User) => !this.friends.includes(user);
      const filterSelf = (user: User) => this.selfUser.id != user.id;

      this.store
        .dispatch("searchByUsername", this.searchQuery)
        .then((users: User[]) => {
          this.users = users.filter(filterFriends).filter(filterSelf);
        });
    },
    addFriend(user: User) {
      this.store.dispatch("sendFriendRequest", user.id);
    },
  },
});
</script>
