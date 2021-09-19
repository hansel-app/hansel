import { Dayjs } from "dayjs";

export const formatDateTime = (date: Dayjs): string => {
  return date.format("D MMM YYYY, h:mm A");
};
