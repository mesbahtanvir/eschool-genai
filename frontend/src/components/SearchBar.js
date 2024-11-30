import React, { useState } from "react";
import axios from "axios";
import { TextField, Button, Box, Container, Typography } from "@mui/material";
import CourseBlueprint from "./CourseBlueprint";

function SearchBar() {
  const [searchQuery, setSearchQuery] = useState("");
  const [result, setResult] = useState("");

  // Retrieve the base URL from the environment variable
  const backendUrl =
    process.env.REACT_APP_BACKEND_URL || "http://localhost:8080"; // Default to localhost if not set

  const handleSearch = async () => {
    try {
      const response = await axios.get(`${backendUrl}/course/generate`, {
        params: { course_hint: searchQuery },
      });
      console.log(response);
      setResult(response.data); // Adjust to match the response structure
    } catch (error) {
      console.error("Error:", error);
      setResult("An error occurred while searching");
    }
  };

  return (
    <Container maxWidth="sm">
      <Box marginY={4}>
        <Typography variant="h5" align="center" gutterBottom>
          Search for a Course
        </Typography>
        <Box display="flex" flexDirection="column" alignItems="center" gap={2}>
          <TextField
            label="Search Query"
            variant="outlined"
            fullWidth
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
          />
          <Button variant="contained" color="primary" onClick={handleSearch}>
            Search
          </Button>
        </Box>
      </Box>
      <Box marginTop={4}>
        {result ? (
          typeof result === "string" ? (
            <Typography color="error" align="center">
              {result}
            </Typography>
          ) : (
            <CourseBlueprint courseId={result.course_id} />
          )
        ) : null}
      </Box>
    </Container>
  );
}

export default SearchBar;
