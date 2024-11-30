import React from "react";
import axios from "axios";
import { Button } from "@mui/material";

const CourseEnroll = ({ courseId }) => {
  const userId = "fake_user_id_123"; // Replace with actual user ID in production

  const handleEnroll = async () => {
    try {
      // Retrieve the base URL from the environment variable
      const backendUrl =
        process.env.REACT_APP_BACKEND_URL || "http://localhost:8080"; // Default to localhost if not set
      const response = await axios.post(`${backendUrl}/course/enroll`, {
        course_id: courseId,
        user_id: userId,
      });
      alert(response.data.message || "Enrolled successfully!");
    } catch (error) {
      console.error("Error enrolling in course:", error);
      alert("Failed to enroll in the course. Please try again.");
    }
  };

  return (
    <Button variant="contained" color="primary" onClick={handleEnroll}>
      Enroll Now
    </Button>
  );
};

export default CourseEnroll;
