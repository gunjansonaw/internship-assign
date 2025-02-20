import React, { useState } from "react";
import "../styles/modal.css";
import { addHoliday } from "../api/holiday"; // Import addHoliday API

const HolidayModal = ({ date, close, refresh }) => {
  const [holidayName, setHolidayName] = useState("");

  const handleSave = async () => {
    try {
      const newHoliday = { date, name: holidayName, month: new Date().getMonth() + 1, year: new Date().getFullYear() };
      await addHoliday(newHoliday);
      refresh(); // Refresh holidays after adding
      close(); // Close modal after saving
    } catch (error) {
      console.error("Error saving holiday", error);
    }
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h2>Add Holiday for {date}</h2>
        <input
          type="text"
          placeholder="Holiday Name"
          value={holidayName}
          onChange={(e) => setHolidayName(e.target.value)}
        />
        <div className="modal-actions">
          <button onClick={handleSave}>Save</button>
          <button onClick={close}>Cancel</button>
        </div>
      </div>
    </div>
  );
};

export default HolidayModal;