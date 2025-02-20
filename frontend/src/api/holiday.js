import axios from "axios";

const API_URL = "http://localhost:5000/api/holidays";

// Fetch holidays based on month and year
export const getHolidays = async (month, year) => {
  try {
    const response = await axios.get(`${API_URL}?month=${month}&year=${year}`);
    return response.data;
  } catch (error) {
    console.error("Error fetching holidays:", error);
    throw error;
  }
};

// Add a new holiday
export const addHoliday = async (holiday) => {
  try {
    const response = await axios.post(API_URL, holiday);
    return response.data;
  } catch (error) {
    console.error("Error adding holiday:", error);
    throw error;
  }
};

// Delete a holiday by ID
export const deleteHoliday = async (id) => {
  try {
    const response = await axios.delete(`${API_URL}/${id}`);
    return response.data;
  } catch (error) {
    console.error("Error deleting holiday:", error);
    throw error;
  }
};
