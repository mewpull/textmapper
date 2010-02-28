package net.sf.lapg.templates.api.impl;

import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.Reader;

import net.sf.lapg.templates.api.IProblemCollector;
import net.sf.lapg.templates.api.IBundleLoader;
import net.sf.lapg.templates.api.TemplatesBundle;

/**
 * Loads templates stored along with java classes (using ClassLoader)
 */
public class ClassTemplateLoader implements IBundleLoader {

	private final ClassLoader loader;
	private final String rootPackage;
	private final String charsetName;

	public ClassTemplateLoader(ClassLoader loader, String rootPackage, String charsetName) {
		this.loader = loader;
		this.rootPackage = rootPackage;
		this.charsetName = charsetName;
	}

	private static String getStreamContents(InputStream stream, String charsetName) {
		StringBuffer contents = new StringBuffer();
		char[] buffer = new char[2048];
		int count;
		try {
			Reader in = new InputStreamReader(stream, charsetName);
			try {
				while ((count = in.read(buffer)) > 0) {
					contents.append(buffer, 0, count);
				}
			} finally {
				in.close();
			}
		} catch (IOException ioe) {
			return null;
		}
		return contents.toString();
	}

	public TemplatesBundle load(String bundleName, IProblemCollector collector) {
		String resourceName = rootPackage + "/" + bundleName.replace('.', '/') + BUNDLE_EXT;
		InputStream s = loader.getResourceAsStream(resourceName);
		if (s == null) {
			return null;
		}
		String name = resourceName.indexOf('/') >= 0 ? resourceName.substring(resourceName.lastIndexOf('/'))
				: resourceName;
		return TemplatesBundle.parse(name, getStreamContents(s, charsetName), bundleName, collector);
	}
}
