import unittest
from unittest.mock import patch

import main


class TestMain(unittest.TestCase):
    def test__div__positive_integers(self):
        result = main.div(6, 2)
        self.assertEqual(result, 3)

    def test__div__negative_integers(self):
        result = main.div(-6, 2)
        self.assertEqual(result, -3)

    def test__div__zero_numerator(self):
        result = main.div(0, 2)
        self.assertEqual(result, 0)

    def test__div__zero_denominator(self):
        with self.assertRaises(ZeroDivisionError):
            result = main.div(6, 0)

    def test__get_user_url(self):
        with patch(
            "main.get_base_url", return_value="www.twitter.com"
        ) as mock_get_base_url:
            url = main.get_user_url("aaronraff_")
            mock_get_base_url.assert_called_once()
            self.assertEqual(url, "https://www.twitter.com/aaronraff_")
