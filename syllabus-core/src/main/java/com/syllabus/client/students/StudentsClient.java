package com.syllabus.client.students;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

@FeignClient(name = "${students.name}")
public interface StudentsClient {

    @GetMapping(value = "${students.path}/{id}/student-profile")
    public StudentResponse getStudentByUserId(@PathVariable(name = "id") String userId);
    
}
