<template>
  <NavBar left-arrow @click-left="goBack" />
  <h2>Add friends</h2>
  <Search placeholder="Search by username" v-model="searchQuery" />
  <CellGroup>
    <FriendCell
      v-for="user in filteredUsers"
      :key="user.id"
      :friend="user"
      :shouldDisplayUsername="true"
    />
  </CellGroup>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { CellGroup, NavBar, Search } from "vant";
import { mockFriends } from "@/interfaces/mockData";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
  components: {
    CellGroup,
    FriendCell,
    NavBar,
    Search,
  },
  data() {
    return {
      mockFriends,
      searchQuery: '',
    };
  },
  computed: {
      filteredUsers() {
          if (!this.searchQuery) {
              // Don't display any users when search bar is empty.
              return [];
          }
          return mockFriends.filter(
              user => user.username.indexOf(this.searchQuery) === 0
          );
      }

  },
  methods: {
    goBack() {
      this.$router.back();
    },
  },
});
</script>
