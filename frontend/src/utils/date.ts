import { Dayjs } from "dayjs";

export const formatDateTime = (date: Dayjs): string => {
  return date.format("D MMMM YYYY h:mm A");
};
