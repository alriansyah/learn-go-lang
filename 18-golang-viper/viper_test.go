package golang_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config") // nama file config
	config.SetConfigType("json") // extensi file
	config.AddConfigPath(".") // path file config

	// read config
	err := config.ReadInConfig() // membaca isi file config
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Al Riansyah", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}
