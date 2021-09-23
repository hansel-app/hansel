<template>
  <div class="background-gradient">
    <Header title="Friends" :backLink="PROFILE_ROUTE" />
    <div class="container">
      <CellGroup>
        <FriendCell
          v-for="friend in sortedFriends"
          :key="friend.id"
          :friend="friend"
          :shouldDisplayUsername="true"
        />
      </CellGroup>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapState, useStore } from "vuex";
import { CellGroup } from "vant";
import { User } from "@/interfaces/user";
import { PROFILE_ROUTE } from "@/constants/routes";
import Header from "@/components/Header.vue";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
  data() {
    return {
      store: useStore(),
      PROFILE_ROUTE,
    };
  },
  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      friends: (state: any) => state.user.friends as User[],
    }),
    // Alphabetically, by displayName
    sortedFriends(): User[] {
      return this.friends.slice().sort((f1, f2) => {
        return f1.displayName.localeCompare(f2.displayName);
      });
    },
  },
  components: {
    FriendCell,
    CellGroup,
    Header,
  },
});
</script>
