<template>
  <van-row :justify="justify" class="fill-width gem-message-closed">
    <van-cell
      v-if="isSentBySelf"
      class="gem-message"
      v-bind:style="gemMessageStyle"
      center
    >
      <template #right-icon>
        <van-icon
          class="gem-icon"
          :name="getGemImageUrl"
          @click="goToDropGem"
        />
      </template>
      <template #title>
        <van-col class="text">
          <p class="content new-gem-text">New gem!</p>
          <p class="address-text">{{ gemLocation }}</p>
        </van-col>
      </template>
    </van-cell>
    <van-cell v-else class="gem-message" v-bind:style="gemMessageStyle" center>
      <template #icon>
        <van-icon
          class="gem-icon"
          :name="getGemImageUrl"
          @click="goToDropGem"
        />
      </template>
      <template #title>
        <van-col class="text">
          <p class="content new-gem-text">New gem!</p>
          <p class="address-text">{{ gemLocation }}</p>
        </van-col>
      </template>
    </van-cell>
  </van-row>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Gem, GemColor } from "@/interfaces";
import { getEnumKeyByEnumValue } from "@/utils/enum";
import { getLocationNameFromLatLng } from "@/utils/geolocation";

export default defineComponent({
  props: {
    gem: {
      type: Object as PropType<Gem>,
      required: true,
    },
    gemMessageStyle: {
      type: String,
      required: true,
    },
    justify: {
      type: String,
      required: true,
    },
    isSentBySelf: {
      type: Boolean,
      required: true,
    },
  },
  mounted() {
    getLocationNameFromLatLng(this.gem.position).then(
      (address) => (this.gemLocation = address)
    );
  },
  data() {
    return {
      gemLocation: "" as string,
    };
  },
  computed: {
    getGemImageUrl(): string {
      const gemColorName = getEnumKeyByEnumValue(GemColor, this.gem.color);
      return require(`@/assets/images/${gemColorName.toLowerCase()}_2048.png`);
    },
  },
});
</script>

<style scoped lang="less">
.gem-message {
  width: 70%;
  padding: 1em 0.5em;
  word-wrap: break-word;
  border-radius: @border-radius-lg;
}

.fill-width {
  width: 100%;
}

.gem-icon {
  font-size: @circle-button-icon-size-sm;
}

.gem-message-closed {
  padding-top: 1em;
  padding-bottom: 1em;
}

.text {
  margin-left: 8px;
  margin-right: 8px;
}
.new-gem-text {
  height: 1rem;
  margin-top: -0.5vh;
  text-align: right;
}

.address-text {
  font-size: @font-size-md;
  height: 2rem;
}
</style>
