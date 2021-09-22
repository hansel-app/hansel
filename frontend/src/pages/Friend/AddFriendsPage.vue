<template>
  <div class="background-gradient">
    <BackSwipe />
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
          <!-- TODO: find correct add user icon -->
          <van-icon name="user-o" @click="addFriend" />
        </FriendCell>
      </CellGroup>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, Search } from "vant";
import { User } from "@/interfaces";
import BackSwipe from "@/components/BackSwipe.vue";
import { useStore } from "vuex";
import FriendCell from "@/components/FriendCell.vue";
import Header from "@/components/Header.vue";

export default defineComponent({
  components: {
    BackSwipe,
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
      this.store.dispatch("searchByUsername", this.searchQuery)
        .then((users: User[]) => {
          this.users = users;
        });
    },
    addFriend(receiverId: number) {
      this.store.dispatch("sendFriendRequest", receiverId);
    },
  },
});
</script>
