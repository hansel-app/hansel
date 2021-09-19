<template>
  <Suspense>
    <template #default>
      <div class="container">
        <gem-map :gems="gems" />
        <div class="profile-button">
          <CircleAvatar
            :avatarUrl="placeholderAvatarUrl"
            :radius="3"
            @click="goToProfile"
            class="profile-button"
          >
            Profile</CircleAvatar
          >
        </div>
      </div>
    </template>
    <template #fallback>
      <!-- TODO: add more useful loading / error message -->
      <div>Loading...</div>
    </template>
  </Suspense>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { PROFILE_ROUTE } from "@/constants";
import GemMap from "@/components/GemMap.vue";
import CircleAvatar from "@/components/CircleAvatar.vue";
import { GemServiceClient } from "@/protobuf/GemServiceClientPb";
import { Gem, GetPendingCollectionByUserRequest } from "@/protobuf/gem_pb";

const SERVER_URL = process.env.VUE_APP_SERVER_URL;
const SERVER_PORT = process.env.VUE_APP_SERVER_PORT;

export default defineComponent({
  components: {
    CircleAvatar,
    GemMap,
  },

  setup() {
    const client = new GemServiceClient(
      `${SERVER_URL}:${SERVER_PORT}`,
      null,
      null
    );
    const request = new GetPendingCollectionByUserRequest();
    request.setUserid(2);

    let gems: Gem[] = [];
    client.getPendingCollectionByUser(request, {}, (err, resp) => {
      // TODO: add error handling
      console.log(err);
      console.log(resp.getGemsList());
      gems = resp.getGemsList();
    });

    return {
      gems,
    };
  },
  methods: {
    goToProfile() {
      this.$router.push(PROFILE_ROUTE);
    },
  },
  data() {
    return {
      // TODO: replace with user's own avatar
      placeholderAvatarUrl:
        "https://media.istockphoto.com/vectors/happy-young-woman-watching-into-rounded-frame-isolated-on-white-3d-vector-id1296058958?b=1&k=20&m=1296058958&s=170667a&w=0&h=6m2FU2hKv6emHjNtdNSqBJR1uMq64smptqwDAZNo6bg=",
    };
  },
});
</script>

<style scoped lang="less">
.container {
  position: relative;
}
.container .profile-button {
  position: absolute;
  z-index: 10;
  top: 4px;
  left: 4px;
}
</style>
