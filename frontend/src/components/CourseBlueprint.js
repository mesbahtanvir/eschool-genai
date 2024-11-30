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
        <Typography
          variant="body1"
          color="textSecondary"
          align="center"
          component
        >
          {description}
        </Typography>
      </Box>

      <Box>
        <Typography variant="h5" align="center" gutterBottom>
          Course Modules
        </Typography>
        <Grid2 container spacing={2}>
          {modules.map((module, i) => (
            <Grid2 item xs={12} key={i}>
              <Paper elevation={2}>
                <Box padding={2}>
                  <Typography variant="h6" align="left">
                    {module.title}
                  </Typography>
                  {module.description && (
                    <Typography
                      variant="body2"
                      color="textSecondary"
                      align="left"
                    >
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
        <CourseEnroll courseId={data.courseId} />
      </Box>
    </Container>
  );
};

export default CourseBlueprint;
