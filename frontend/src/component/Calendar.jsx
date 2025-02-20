import React, { useState, useEffect, useCallback } from "react";
import "../styles/calendar.css";
import HolidayModal from "./HolidayModal";
import { getHolidays, deleteHoliday } from "../api/holiday"; // Import deleteHoliday API

const Calendar = () => {
  const [currentDate, setCurrentDate] = useState(new Date());
  const [holidays, setHolidays] = useState([]);
  const [modalDate, setModalDate] = useState(null);

  const startYear = 2024;
  const endYear = 2029;

  // Fetch holidays when the month/year changes
  const fetchHolidays = useCallback(async () => {
    try {
      const res = await getHolidays(currentDate.getMonth() + 1, currentDate.getFullYear());
      setHolidays(res); // Update holidays state with fetched data
    } catch (error) {
      console.error("Error fetching holidays", error);
    }
  }, [currentDate]);

  useEffect(() => {
    fetchHolidays();
  }, [fetchHolidays]);

  // Handle the previous month button
  const handlePrevMonth = () => {
    const newDate = new Date(currentDate.getFullYear(), currentDate.getMonth() - 1);
    if (newDate.getFullYear() >= startYear) {
      setCurrentDate(newDate);
    }
  };

  // Handle the next month button
  const handleNextMonth = () => {
    const newDate = new Date(currentDate.getFullYear(), currentDate.getMonth() + 1);
    if (newDate.getFullYear() <= endYear) {
      setCurrentDate(newDate);
    }
  };

  // Handle deleting a holiday
  const handleDeleteHoliday = async (id) => {
    try {
      await deleteHoliday(id);
      fetchHolidays(); // Refresh holidays after deleting
    } catch (error) {
      console.error("Error deleting holiday", error);
    }
  };

  const daysInMonth = new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 0).getDate();
  const firstDay = new Date(currentDate.getFullYear(), currentDate.getMonth(), 1).getDay();

  return (
    <div className="calendar-container">
      <div className="calendar-header">
        <button onClick={handlePrevMonth} disabled={currentDate.getFullYear() <= startYear && currentDate.getMonth() <= 0}>◀</button>
        <h2>
          {currentDate.toLocaleString("default", { month: "long" })} {currentDate.getFullYear()}
        </h2>
        <button onClick={handleNextMonth} disabled={currentDate.getFullYear() >= endYear && currentDate.getMonth() >= 11}>▶</button>
      </div>

      <div className="calendar-grid">
        {["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"].map((day) => (
          <div key={day} className="calendar-day">{day}</div>
        ))}

        {Array(firstDay).fill("").map((_, i) => <div key={i} className="empty"></div>)}

        {Array.from({ length: daysInMonth }, (_, i) => {
          const day = i + 1;
          const holiday = holidays.find((h) => h.date === day);

          return (
            <div key={day} className="calendar-cell" onClick={() => setModalDate(day)}>
              <span>{day}</span>
              {holiday && (
                <div className="holiday-label">
                  {holiday.name}
                  <button onClick={(e) => { e.stopPropagation(); handleDeleteHoliday(holiday.id); }} className="delete-button">
                    🗑️
                  </button>
                </div>
              )}
            </div>
          );
        })}
      </div>

      {modalDate && <HolidayModal date={modalDate} close={() => setModalDate(null)} refresh={fetchHolidays} />}
    </div>
  );
};

export default Calendar;
