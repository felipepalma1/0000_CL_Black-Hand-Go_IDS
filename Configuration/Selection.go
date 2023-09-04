package Configuration

import "0000_CL_Black-Hand-Go_IDS/Exercise"

func Selection(selectedActivity Activity) {
	switch selectedActivity.ID {
	case "03":
		Exercise.E001()
	}
}
