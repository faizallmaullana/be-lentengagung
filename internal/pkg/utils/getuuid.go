package utils

// func GetUserIDFromContext(c *gin.Context) (string, error) {
// 	userIDValue, exists := c.Get("id_user")
// 	if !exists {
// 		return uuid.Nil, fmt.Errorf("user_id not found in context")
// 	}

// 	switch v := userIDValue.(type) {
// 	case string:
// 		id, err := uuid.Parse(v)
// 		if err != nil {
// 			return uuid.Nil, fmt.Errorf("invalid uuid string in context: %w", err)
// 		}
// 		return id, nil
// 	case string:
// 		return v, nil
// 	default:
// 		return uuid.Nil, fmt.Errorf("unsupported id_user type in context: %T", userIDValue)
// 	}
// }
