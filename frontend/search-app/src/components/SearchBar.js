import React, { useState } from "react";
import axios from "axios";
import "./SearchBar.css";
import CourseBlueprint from "./CourseBlueprint";

function SearchBar() {
  const [searchQuery, setSearchQuery] = useState("");
  const [result, setResult] = useState("");

  const handleSearch = async () => {
    try {
      const response = await axios.get(
        "http://127.0.0.1:5000/generate_course",
        {
          params: { input_text: searchQuery },
        }
      );
      setResult(response.data.output); // Adjust to match the response structure
    } catch (error) {
      console.error("Error:", error);
      setResult("An error occurred while searching");
    }
  };

  return (
    <div className="search-container">
      <input
        type="text"
        value={searchQuery}
        onChange={(e) => setSearchQuery(e.target.value)}
        placeholder="Enter your search query..."
        className="search-input"
      />
      <button onClick={handleSearch} className="search-button">
        Search
      </button>
      {result && <CourseBlueprint data={result} />}
    </div>
  );
}

export default SearchBar;
