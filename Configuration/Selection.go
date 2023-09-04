package Configuration

import "0000_CL_Black-Hand-Go_IDS/Exercise"

func Selection(selectedActivity Activity) {
	switch selectedActivity.ID {
	case "01":
		Exercise.E001()

	case "02":
		Exercise.E002()
	}
}
