import React, { useState } from "react";
import axios from "axios";
import "./SearchBar.css";

function SearchBar() {
  const [searchQuery, setSearchQuery] = useState("");
  const [result, setResult] = useState("");

  const handleSearch = async () => {
    try {
      const response = await axios.post(
        "http://127.0.0.1:5000/api/generate_course",
        {
          prompt: searchQuery,
        }
      );
      setResult(response.data.response);
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
      {result && <pre className="result">{result}</pre>}
    </div>
  );
}

export default SearchBar;
