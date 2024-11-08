import React from "react";

export const CourseBlueprint = ({ data }) => {
  console.log(data);
  if (!data || !data.modules) {
    return <p>Loading course data...</p>; // Add a fallback for undefined data
  }
  return (
    <div style={styles.container}>
      <h1 style={styles.title}>{data.title}</h1>
      <p style={styles.overview}>{data.overview}</p>
      <h2 style={styles.objectiveTitle}>Course Objective</h2>
      <p style={styles.objective}>{data.objective}</p>

      <h2 style={styles.moduleTitle}>Course Modules</h2>
      <div style={styles.modules}>
        {data.modules.map((module, index) => (
          <div key={index} style={styles.module}>
            <h3 style={styles.moduleTitle}>{module.title}</h3>
            <ul style={styles.topicList}>
              {module.topics.map((topic, i) => (
                <li key={i} style={styles.topicItem}>
                  {topic}
                </li>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </div>
  );
};

const styles = {
  container: {
    maxWidth: "800px",
    margin: "auto",
    padding: "20px",
    fontFamily: "Arial, sans-serif",
    color: "#333",
    backgroundColor: "#f9f9f9",
    borderRadius: "8px",
    boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)",
  },
  title: {
    fontSize: "2rem",
    color: "#333",
    textAlign: "center",
    marginBottom: "10px",
  },
  overview: {
    fontSize: "1.1rem",
    color: "#555",
    marginBottom: "20px",
    lineHeight: "1.5",
  },
  objectiveTitle: {
    fontSize: "1.5rem",
    color: "#222",
    marginTop: "20px",
    marginBottom: "5px",
  },
  objective: {
    fontSize: "1rem",
    color: "#555",
    marginBottom: "20px",
  },
  moduleTitle: {
    fontSize: "1.5rem",
    color: "#444",
    marginTop: "20px",
    marginBottom: "10px",
  },
  modules: {
    marginTop: "10px",
  },
  module: {
    backgroundColor: "#fff",
    borderRadius: "8px",
    padding: "15px",
    marginBottom: "15px",
    boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
  },
  topicList: {
    paddingLeft: "20px",
    margin: 0,
  },
  topicItem: {
    fontSize: "1rem",
    color: "#333",
    lineHeight: "1.5",
    marginBottom: "8px",
  },
};

export default CourseBlueprint;
