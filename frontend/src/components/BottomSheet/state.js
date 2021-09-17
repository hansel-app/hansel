import { ref, computed } from "vue"

export const collapsed = ref(true)
export const toggleBottomSheet = () => (collapsed.value = !collapsed.value)

export const BOTTOM_SHEET_HEIGHT = 180
export const BOTTOM_SHEET_HEIGHT_COLLAPSED = 38
export const bottomSheetHeight = computed(
    () => `${collapsed.value ? BOTTOM_SHEET_HEIGHT_COLLAPSED : BOTTOM_SHEET_HEIGHT}px`
)