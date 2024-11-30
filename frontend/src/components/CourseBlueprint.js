import React from "react";
import { Container, Typography, Box, Paper, Grid2 } from "@mui/material";
import CourseEnroll from "./CourseEnroll";

export const CourseBlueprint = ({ data }) => {
  if (!data?.course?.course_blueprint)
    return <Typography>Loading course data...</Typography>;

  const { title, description, modules } = data.course.course_blueprint;

  return (
    <Container maxWidth="md">
      <Box>
        <Typography variant="h4" align="center" gutterBottom>
          {title}
        </Typography>
        <Typography variant="body1" color="textSecondary" paragraph>
          {description}
        </Typography>
      </Box>

      <Box>
        <Typography variant="h5" gutterBottom>
          Course Modules
        </Typography>
        <Grid2 container spacing={2}>
          {modules.map((module, i) => (
            <Grid2 item xs={12} key={i}>
              <Paper elevation={2}>
                <Box padding={2}>
                  <Typography variant="h6">{module.title}</Typography>
                  {module.description && (
                    <Typography variant="body2" color="textSecondary">
                      {module.description}
                    </Typography>
                  )}
                </Box>
              </Paper>
            </Grid2>
          ))}
        </Grid2>
      </Box>

      <Box textAlign="center" marginTop={4}>
        <CourseEnroll />
      </Box>
    </Container>
  );
};

export default CourseBlueprint;
