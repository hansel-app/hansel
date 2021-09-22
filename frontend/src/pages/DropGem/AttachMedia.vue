<template>
  <div>
    <GemHeader title="Drop a gem" :color="color" />
    <GemColorCarousel @selected-color-changed="onSelectedColorChanged" />
    <div class="container">
      <p class="sub-header">Choose media</p>

      <Row justify="start">
        <Uploader
          v-model="fileList"
          :max-count="1"
          :preview-size="150"
          :after-read="afterFileRead"
        >
        </Uploader>
      </Row>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import GemHeader from "@/components/GemHeader.vue";
import GemColorCarousel from "./GemColorCarousel.vue";
import { GemColor } from "@/interfaces";
import { mapState, useStore } from "vuex";
import { Uploader, Row, UploaderFileListItem } from "vant";

export default defineComponent({
  components: {
    GemHeader,
    Uploader,
    GemColorCarousel,
    Row,
  },
  setup() {
    const store = useStore();
    return { store };
  },
  data() {
    return {
      fileList: [] as UploaderFileListItem[],
    };
  },
  computed: {
    ...mapState({
      color: (state: any) => state.gems.dropGemFormState.color as GemColor,
    }),
  },
  methods: {
    onSelectedColorChanged(selectedColor: GemColor) {
      this.$emit("set-gem-color-event", selectedColor);
    },
    afterFileRead(fileListItem: UploaderFileListItem) {
      this.store.commit("updateDropGemFormState", {
        attachment: fileListItem.file,
      });
    },
  },
});
</script>

<style scoped lang="less"></style>
