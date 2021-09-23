<template>
  <div>
    <Header :isCloseWindow="true" title="Drop a gem" />
    <div class="container">
      <p class="sub-header">Choose a receiver</p>
      <Searchbar placeholder="Search a friend" />
      <CellGroup>
        <FriendCell
          v-for="friend in filteredFriends"
          :key="friend.userId"
          :friend="friend"
          :isClickable="true"
          :shouldDisplayUsername="true"
          @click="
            () => {
              $emit('set-receiver-event', friend.userId);
              $emit('next-stage');
            }
          "
        />
      </CellGroup>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { CellGroup } from "vant";
import { User } from "@/interfaces";
import FriendCell from "@/components/FriendCell.vue";
import Searchbar from "@/components/Searchbar.vue";
import Header from "@/components/Header.vue";
import { mapState, useStore } from "vuex";

export default defineComponent({
  components: {
    CellGroup,
    FriendCell,
    Header,
    Searchbar,
  },
  setup() {
    const store = useStore();
    const fetchFriends = () => store.dispatch("getFriends");

    onMounted(() => {
      fetchFriends();
    });
  },
  data() {
    return {
      searchQuery: "",
    };
  },
  computed: {
    ...mapState({
      allFriends: (state: any) => state.user.friends as User[],
    }),
    filteredFriends(): User[] {
      if (this.searchQuery.length === 0) {
        return this.allFriends;
      }
      return this.allFriends.filter((friend) => {
        const matchUsername = friend.username
          .toLowerCase()
          .includes(this.searchQuery.toLowerCase());
        const matchDisplayname = friend.displayName
          .toLowerCase()
          .includes(this.searchQuery.toLowerCase());
        return matchUsername || matchDisplayname;
      });
    },
  },
});
</script>
<style scoped>
.van-cell-group {
  background: none;
}
</style>
