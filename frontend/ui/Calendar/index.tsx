"use client";

import { useState } from "react";
import ReactCalendar from "react-calendar";
import { differenceInCalendarDays, addDays } from "date-fns";
import { DotFilledIcon } from "@radix-ui/react-icons";

function isSameDay(a: Date, b: Date) {
  return differenceInCalendarDays(a, b) === 0;
}

const datesToAddContentTo = [
  new Date(),
  addDays(new Date(), 1),
  addDays(new Date(), 5),
];

function tileContent({ date, view }: any) {
  // Add class to tiles in month view only
  if (view === "month") {
    // Check if a date React-Calendar wants to check is on the list of dates to add class to
    if (datesToAddContentTo.find((dDate) => isSameDay(dDate, date))) {
      return (
        <DotFilledIcon
          transform="scale(1.2)"
          style={{
            alignSelf: "start",
            position: "absolute",
            top: "3px",
            right: "1px",
          }}
        />
      );
    }
  }
}

export function Calendar() {
  const [date, setDate] = useState(new Date());

  return (
    <>
      <ReactCalendar
        onChange={setDate}
        value={date}
        minDate={new Date(2022, 1, 1)}
        tileContent={tileContent}
      />
    </>
  );
}
