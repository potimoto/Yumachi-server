CREATE TABLE `pairing_user` (
  `app_id` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `pairing_id` varchar(5) COLLATE utf8_unicode_ci DEFAULT NULL,
  `gender` varchar(5) COLLATE utf8_unicode_ci DEFAULT NULL,
  `beacon_id` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `mark_str` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `RasPi_id` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
