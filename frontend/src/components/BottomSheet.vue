<template>
  <div 
    class="drawer" 
    @click="toggleBottomSheet" 
    v-touch:swipe.top="swipeBottomSheetUp"
    v-touch:swipe.bottom="swipeBottomSheetDown"
  >
    <div class="wrapper">
      <van-row ref="infoBox" class="message-content">
        <p>{{ gem.message }}</p>
      </van-row>
      <div v-if="!collapsed">
        <FriendCell
        :key="friend.id"
        :friend="friend"
        :isClickable="true"
        :shouldDisplayUsername="true"
        />
        <van-row class="message-details content">
          <img :src="MapIcon" id="icon">
          <van-col span="19">
            <p>{{ gemAddress }}</p>
          </van-col>
        </van-row>
        <van-row class="message-details content">
          <img :src="CalendarIcon" id="icon">
          <van-col span="19">
            <p>{{ displayDropDateTime }}</p>
          </van-col>
        </van-row>
        <van-button @click="backToHome">Back to map!</van-button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Gem, User } from "@/interfaces";
import { defineComponent, PropType, ref, computed } from "vue";
import { Row, Button } from "vant";
import { formatDateTime } from "@/utils/date";
import { HOME_ROUTE } from "@/constants";
import CalendarIcon from "@/assets/icons/calendar.svg";
import MapIcon from "@/assets/icons/map.svg";
import FriendCell from "@/components/FriendCell.vue"
import axios from "axios";

const GOOGLE_API_KEY = window.env.VUE_APP_GOOGLE_API_KEY;

export default defineComponent({
  setup() {
    const collapsed = ref(true);
    const toggleBottomSheet = () => (collapsed.value = !collapsed.value);
    // how to set this dynamically...
    const BOTTOM_SHEET_HEIGHT = 160;
    const BOTTOM_SHEET_HEIGHT_COLLAPSED = 50;
    const bottomSheetHeight = computed(
      () =>
        `${
          collapsed.value ? BOTTOM_SHEET_HEIGHT_COLLAPSED : BOTTOM_SHEET_HEIGHT
        }px`
    );

    return {
      collapsed,
      toggleBottomSheet,
      bottomSheetHeight,
      CalendarIcon,
      MapIcon,
    };
  },
  mounted () {
    axios
      .get(
        'https://maps.googleapis.com/maps/api/geocode/json?latlng=' + this.gem.position.lat + ',' + this.gem.position.lng + '&key=' + GOOGLE_API_KEY)
      .then(response => (this.gemAddress = response.data.results[1].formatted_address));
  },
  components: {
    "van-button": Button,
    "van-row": Row,
    FriendCell,
  },
  props: {
    gem: {
      type: Object as PropType<Gem>,
      required: true,
    }
  },
  data() {
    return { 
      friend: this.gem.createdBy as User,
      gemAddress: null,
    };
  },
  computed: {
    displayDropDateTime() {
      return `${formatDateTime(this.gem.createdAt)}`;
    },
  },
  methods: {
    backToHome() {
      this.$router.replace(HOME_ROUTE);
    },
    swipeBottomSheetDown() {
      if (this.collapsed == false) {
        this.collapsed = !this.collapsed;
      }
    },
    swipeBottomSheetUp() {
      if (this.collapsed == true) {
        this.collapsed = !this.collapsed;
      }
    },
  },
});
</script>

<style scoped>
.drawer {
  color: black;
  background-color: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(5px);
  border-radius: 2em 2em 0 0;

  float: left;
  position: fixed;
  z-index: 1;
  left: 0;
  right: 0;
  bottom: 0;
  padding: 0.5em;

  transition: max-height 0.3s;

  display: flex;
  flex-direction: row;
}

.wrapper {
  float: left;
  width: 100%;
  text-align: left;
  padding: 0 2em;
}

#icon {
  padding-right: 1em;
  float: left;
}

.van-button {
  margin: 2em 0;
  padding: 1em 2em;
  background-color: black;
  border: none;
  border-radius: 2em;
  color: white;
  margin-left: 50%;
}

.van-col {
  overflow: hidden;
  overflow-wrap: break-all;
}

.image-wrapper {
  clear: both;
}
</style>
