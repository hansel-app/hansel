<template>
  <div>
    <div class="overlay remove-margin">
      <Header
        justify="space-between"
        class="log-details-header"
        title=""
        :backLink="GEM_LOGS_ROUTE"
      >
        <van-col class="friend-name-container">
          <van-row class="friend-displayname">
            <p>{{ friend.displayName }}</p>
          </van-row>
          <van-row class="friend-username">
            <p>@{{ friend.username }}</p>
          </van-row>
        </van-col>
        <van-col>
          <CircleAvatar
            class="friend-avatar"
            :avatarUrl="friend.avatar"
            :radius="2.5"
          />
        </van-col>
      </Header>

      <div class="gem-messages-container">
        <GemMessage
          v-for="gemAndIsSentBySelf in gemsAndIsSentBySelf"
          :key="gemAndIsSentBySelf.gem.id"
          :gem="gemAndIsSentBySelf.gem"
          :isSentBySelf="gemAndIsSentBySelf.isSentBySelf"
        />
      </div>
    </div>
    <div class="reverse-background-gradient">
      <div class="base">
        <div class="header" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { GEM_LOGS_ROUTE } from "@/constants";
import CircleAvatar from "@/components/CircleAvatar.vue";
import Header from "@/components/Header.vue";
import GemMessage from "./GemMessage.vue";
import { mapState, useStore } from "vuex";
import { Gem, GemLogsWithFriend, User } from "@/interfaces";

export default defineComponent({
  components: {
    Header,
    CircleAvatar,
    GemMessage,
  },

  setup() {
    const store = useStore();

    return {
      store,
      GEM_LOGS_ROUTE,
    };
  },
  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      gemLog: (state: any) => {
        const logs = state.gems.selectedGemLog as GemLogsWithFriend;
        const sortedGems = [...logs.gems].sort((gem1, gem2) =>
          gem1.createdAt.diff(gem2.createdAt)
        );

        return {
          friend: logs.friend,
          gems: sortedGems,
        } as GemLogsWithFriend;
      },
      selfUser: (state: any) => state.user.self as User,
    }),
    gemsAndIsSentBySelf(): { gem: Gem; isSentBySelf: boolean }[] {
      return this.gemLog.gems.map((gem) => {
        return {
          gem: gem,
          isSentBySelf: gem.createdBy.userId === this.selfUser.userId,
        };
      });
    },
    friend(): User {
      console.assert(
        this.gemLog.friend !== undefined,
        "Friend must be associated with a gem log"
      );
      return this.gemLog.friend!;
    },
  },
});
</script>

<style scoped lang="less">
.friend-name-container {
  width: 50%;
  height: 50%;
  .friend-displayname {
    color: @text-color;
    font-size: @font-size-xl;
    height: 3rem;
  }
  .friend-username {
    color: @gray-6;
    font-size: @font-size-md;
    height: 3rem;
  }
}

.gem-messages-container {
  margin-left: 10vw;
  margin-right: 10vw;
}

.header {
  height: 16vh;
  background-color: white;
}

.log-details-header {
  height: 12vh;
  max-height: 12vh;
  width: 90%;
}

.remove-margin {
  margin-top: -5vh;
}
</style>
