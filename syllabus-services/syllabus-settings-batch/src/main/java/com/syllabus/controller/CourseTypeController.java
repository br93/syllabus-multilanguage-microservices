package com.syllabus.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import com.syllabus.cache.CacheService;
import com.syllabus.helper.CsvHelper;
import com.syllabus.message.ResponseMessage;
import com.syllabus.service.impl.CourseTypeService;

import lombok.RequiredArgsConstructor;

@RequiredArgsConstructor
@RestController
@RequestMapping("api/v1/csv")
public class CourseTypeController {

	private final CourseTypeService courseTypeService;
	private final CacheService cacheService;

	@PostMapping("course-types")
	public ResponseEntity<ResponseMessage> uploadFile(@RequestParam("file") MultipartFile file) {
		String message = "";

		if (CsvHelper.hasCSVFormat(file)) {
			try {
				courseTypeService.save(file);
				message = "Uploaded the file successfully: " + file.getOriginalFilename();
				cacheService.flush();
				return ResponseEntity.status(HttpStatus.CREATED).body(new ResponseMessage(message));
			} catch (Exception ex) {
				message = "Could not upload the file: " + file.getOriginalFilename();
				return ResponseEntity.status(HttpStatus.EXPECTATION_FAILED).body(new ResponseMessage(message));
			}

		}

		message = "Please upload a csv file";
		return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(new ResponseMessage(message));
	}

}
