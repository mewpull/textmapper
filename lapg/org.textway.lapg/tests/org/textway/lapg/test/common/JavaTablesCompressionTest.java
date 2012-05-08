/**
 * Copyright 2002-2012 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textway.lapg.test.common;

import org.junit.Test;
import org.textway.lapg.common.JavaArrayArchiver;
import org.textway.lapg.gen.TemplateStaticMethods;
import org.textway.lapg.test.gen.JavaTemplateRoutines;

import java.util.Arrays;
import java.util.List;

import static org.junit.Assert.*;

public class JavaTablesCompressionTest {

	private void checkDecompression(int[][] a) {
		String c = JavaArrayArchiver.packIntInt(a, 5);
		String starts = a.length + "," + a[0].length + ",\n";
		assertTrue(c.startsWith(starts));
		c = c.substring(starts.length());

		StringBuilder extractedString = new StringBuilder();
		char[] chs = c.toCharArray();
		boolean isstring = false;
		for (char ch : chs) {
			if (ch == '"') {
				isstring = !isstring;
				continue;
			}
			if (isstring) {
				extractedString.append(ch);
			}
		}

		int[][] b = JavaArrayArchiver.unpackIntInt(a.length, a[0].length, extractedString.toString());
		for (int i = 0; i < a.length; i++) {
			for (int e = 0; e < a[0].length; e++) {
				if (a[i][e] != b[i][e]) {
					fail("wrong decompression at " + i + "," + e);
				}
			}
		}
	}

	private void checkDecompression(int[] a) {
		List<List<String>> packed = TemplateStaticMethods.packInt(a);
		String parsed = parsePackedJavaString(packed);
		int[] b = JavaTemplateRoutines.test_unpack_int(a.length, parsed);
		assertArrayEquals(a, b);
	}

	private void checkDecompression(short[] a) {
		List<List<String>> packed = TemplateStaticMethods.packShort(a);
		String parsed = parsePackedJavaString(packed);
		short[] b = JavaTemplateRoutines.test_unpack_short(a.length, parsed);
		assertArrayEquals(a, b);
	}

	private String parsePackedJavaString(List<List<String>> c) {
		StringBuilder extractedString = new StringBuilder();
		for (List<String> slist : c) {
			for (String s : slist) {
				parseJavaString(extractedString, s);
			}
		}
		return extractedString.toString();
	}

	private void parseJavaString(StringBuilder result, String javaStringInQuotes) {
		assertEquals('"', javaStringInQuotes.charAt(0));
		assertEquals('"', javaStringInQuotes.charAt(javaStringInQuotes.length() - 1));

		char[] chs = javaStringInQuotes.substring(1, javaStringInQuotes.length() - 1).toCharArray();
		int i = 0;
		while (i < chs.length) {
			assertEquals('\\', chs[i++]);
			if (chs[i] == 'u') {
				try {
					result.append((char) Integer.parseInt(new String(chs, i + 1, 4), 16));
				} catch (NumberFormatException ex) {
					fail(ex.toString());
				}
				i += 5;
			} else {
				int start = i;
				while (i < chs.length && chs[i] >= '0' && chs[i] <= '9') {
					i++;
				}
				assertTrue(new String(chs, start, chs.length - start), i > start);
				try {
					result.append((char) Integer.parseInt(new String(chs, start, i - start), 8));
				} catch (NumberFormatException ex) {
					fail(ex.toString());
				}
			}
		}
	}

	// TODO FIXME test short compression

	@Test
	public void testCompression1() {
		checkDecompression(new int[][]{
				{1, 2},
				{3, 3}
		});
	}

	@Test
	public void testCompression2() {
		checkDecompression(new int[][]{
				{1, 1},
				{1, 1}
		});
	}

	@Test
	public void testCompression3() {
		checkDecompression(new int[][]{
				{0, 0},
				{0, 1}
		});
	}

	@Test
	public void testCompression4() {
		checkDecompression(new int[][]{
				{0},
				{0}
		});
	}

	@Test
	public void testCompression5() {
		checkDecompression(new int[][]{
				{-789}
		});
	}

	@Test
	public void testCompression6() {
		checkDecompression(new int[][]{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{3, 4, 5, 7, 8, 8, 8, 8}
		});
	}

	@Test
	public void testCompressionBig() {
		checkDecompression(new int[][]{
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
				{1, 2, 3, 4, 5, 6, 7, 8, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 2, 1, 1, 1, 1, 1, 1, 43, 345, 345, 345, 34, 34, 345, 34, 345, 34533, 333},
		});
	}

	@Test
	public void testIntCompression1() {
		checkDecompression(new int[]{1, 2, 3, 4, 5, 6, 7, 8});
	}

	@Test
	public void testIntCompressionEmpty() {
		checkDecompression(new int[]{});
	}

	@Test
	public void testIntCompressionOne() {
		checkDecompression(new int[]{-100});
		checkDecompression(new int[]{1});
	}

	@Test
	public void testIntCompressionNearZero() {
		checkDecompression(new int[]{0});
		checkDecompression(new int[]{-1});
	}

	@Test
	public void testIntCompression2() {
		checkDecompression(new int[]{0, 0, 0});
		checkDecompression(new int[]{Integer.MAX_VALUE});
	}

	@Test
	public void testIntCompressionMax() {
		checkDecompression(new int[]{Integer.MAX_VALUE});
	}

	@Test
	public void testIntCompressionMin() {
		checkDecompression(new int[]{Integer.MIN_VALUE});
	}

	@Test
	public void testIntCompressionLong1() {
		int[] s = new int[32768];
		Arrays.fill(s, -1);
		checkDecompression(s);
	}

	@Test
	public void testIntCompressionLong2() {
		int[] s = new int[65536];
		for(int i = 0; i < s.length; i++) {
			s[i] = i - 32768;
		}
		checkDecompression(s);
	}

	@Test
	public void testShortCompression1() {
		checkDecompression(new short[]{1, 2, 3, 4, 5, 6, 7, 8});
	}

	@Test
	public void testShortCompressionEmpty() {
		checkDecompression(new short[]{});
	}

	@Test
	public void testShortCompressionOne() {
		checkDecompression(new short[]{-100});
		checkDecompression(new short[]{1});
	}

	@Test
	public void testShortCompressionNearZero() {
		checkDecompression(new short[]{0});
		checkDecompression(new short[]{-1});
	}

	@Test
	public void testShortCompression2() {
		checkDecompression(new short[]{0, 0, 0});
		checkDecompression(new short[]{Short.MAX_VALUE});
	}

	@Test
	public void testShortCompressionMax() {
		checkDecompression(new short[]{Short.MAX_VALUE});
	}

	@Test
	public void testShortCompressionMin() {
		checkDecompression(new short[]{Short.MIN_VALUE});
	}

	@Test
	public void testShortCompressionLong() {
		short[] s = new short[32768*4 + 7];
		Arrays.fill(s, (short) -1);
		checkDecompression(s);
	}
}
