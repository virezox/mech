package com.bandcamp.android.network;
public class a {
    private static final com.bandcamp.shared.util.BCLog a;

    static a()
    {
        com.bandcamp.android.network.a.a = com.bandcamp.shared.util.BCLog.d;
        return;
    }

    public a()
    {
        com.bandcamp.shared.network.API.a("134", com.bandcamp.android.network.a.b());
        return;
    }

    private static String a(String p8, int p9)
    {
        char[] v8_1 = p8.toCharArray();
        int v0 = v8_1.length;
        int v2 = 0;
        if ((p9 % 2) != 0) {
            int v1_1 = 3;
            while ((v1_1 * v1_1) <= p9) {
                if ((p9 % v1_1) != 0) {
                    v1_1 += 2;
                } else {
                    int v1_0 = 0;
                }
                if (v1_0 != 0) {
                    p9++;
                }
            }
            v1_0 = 1;
        }
        int v1_3 = new boolean[(p9 + 1)];
        v1_3[0] = 1;
        v1_3[1] = 1;
        int v4_1 = -1;
        int v5_0 = -1;
        int v6 = 2;
        while (v6 <= p9) {
            if (v1_3[v6] == 0) {
                int v4_9 = (v6 + v6);
                while (v4_9 <= p9) {
                    v1_3[v4_9] = 1;
                    v4_9 += v6;
                }
                v4_1 = v5_0;
                v5_0 = v6;
            }
            v6++;
        }
        String v9_1 = (10 - (v4_1 % 10));
        int v1_4 = (26 - (v5_0 % 26));
        while (v2 < v0) {
            int v3_0;
            v3_0 = v8_1[v2];
            if ((v3_0 < 97) || (v3_0 > 122)) {
                if ((v3_0 < 65) || (v3_0 > 90)) {
                    if ((v3_0 >= 48) && (v3_0 <= 57)) {
                        v3_0 += v9_1;
                        if (v3_0 > 57) {
                            v3_0 -= 10;
                        }
                    }
                } else {
                    v3_0 += v1_4;
                    if (v3_0 > 90) {
                        v3_0 -= 26;
                    }
                }
            } else {
                v3_0 += v1_4;
                if (v3_0 > 122) {
                }
            }
            v8_1[v2] = ((char) v3_0);
            v2++;
        }
        return new String(v8_1);
    }

    private java.net.URL a(String p3, java.util.List p4)
    {
        String v4_1 = com.bandcamp.shared.network.d.a(p4);
        if (!android.text.TextUtils.isEmpty(v4_1)) {
            com.bandcamp.shared.network.API v0_3 = new StringBuilder();
            v0_3.append("?");
            v0_3.append(v4_1);
            v4_1 = v0_3.toString();
        }
        com.bandcamp.shared.network.API v0_0 = com.bandcamp.shared.network.API.a();
        StringBuilder v1_1 = new StringBuilder();
        v1_1.append(p3);
        v1_1.append(v4_1);
        return v0_0.f(v1_1.toString());
    }

    private static String b()
    {
        return com.bandcamp.android.network.a.a("0vhT01EnLU2mFu8x/wlE1EhDDkXSdWYEST5kIISgQet=", 216343);
    }

    public com.bandcamp.android.util.Promise a()
    {
        com.bandcamp.android.util.Promise v0_1 = new com.bandcamp.android.network.f();
        String v1_3 = com.bandcamp.shared.platform.e.a().e();
        v0_1.a("platform", "android");
        int v2_0 = new StringBuilder();
        v2_0.append(v1_3.b);
        v2_0.append("-");
        v2_0.append(v1_3.a);
        v0_1.a("version", v2_0.toString());
        return this.a("/api/mobile/24/serverinfo", 0, v0_1);
    }

    public com.bandcamp.android.util.Promise a(char p3, long p4)
    {
        com.bandcamp.android.network.f v0_1 = new com.bandcamp.android.network.f();
        v0_1.a("tralbum_type", Character.toString(p3));
        v0_1.a("tralbum_id", Long.toString(p4));
        return this.a("/api/mobile/24/tralbum_lyrics", 0, v0_1);
    }

    public com.bandcamp.android.util.Promise a(char p3, long p4, long p6)
    {
        java.util.HashMap v0_1 = new java.util.HashMap();
        v0_1.put("tralbum_type", Character.valueOf(p3));
        v0_1.put("tralbum_id", Long.valueOf(p4));
        v0_1.put("band_id", Long.valueOf(p6));
        this.a(v0_1);
        return this.a("/api/mobile/24/wishlist_remove", 1, v0_1);
    }

    public com.bandcamp.android.util.Promise a(char p3, long p4, long p6, String p8)
    {
        java.util.HashMap v0_1 = new java.util.HashMap();
        v0_1.put("tralbum_type", Character.valueOf(p3));
        v0_1.put("tralbum_id", Long.valueOf(p4));
        v0_1.put("band_id", Long.valueOf(p6));
        if ((p8 == null) || (p8.isEmpty())) {
            Object[] v5 = new Object[1];
            v5[0] = "API: wishlist action without from parameter, please fix";
            com.bandcamp.android.network.a.a.d(v5);
        } else {
            v0_1.put("from", p8);
        }
        this.a(v0_1);
        return this.a("/api/mobile/24/wishlist_add", 1, v0_1);
    }

    public com.bandcamp.android.util.Promise a(long p2, char p4, long p5)
    {
        com.bandcamp.android.network.f v0_1 = new com.bandcamp.android.network.f();
        v0_1.a("band_id", Long.valueOf(p2));
        v0_1.a("tralbum_type", Character.toString(p4));
        v0_1.a("tralbum_id", Long.valueOf(p5));
        return this.a("/api/mobile/24/tralbum_tags", 0, v0_1);
    }

    public com.bandcamp.android.util.Promise a(String p2, boolean p3)
    {
        com.bandcamp.android.util.Promise v2_1 = com.bandcamp.shared.network.API.a().g(p2);
        v2_1.b(p3);
        v2_1.b("X-Requested-With", "com.bandcamp.android");
        return com.bandcamp.android.network.c.a(v2_1);
    }

    public com.bandcamp.android.util.Promise a(String p1, boolean p2, java.util.Collection p3)
    {
        return com.bandcamp.android.network.c.a(this.b(p1, p2, p3));
    }

    public com.bandcamp.android.util.Promise a(String p2, boolean p3, java.util.Map p4)
    {
        return this.a(p2, p3, new org.json.JSONObject(p4));
    }

    public com.bandcamp.android.util.Promise a(String p1, boolean p2, org.json.JSONObject p3)
    {
        return com.bandcamp.android.network.c.a(this.b(p1, p2, p3));
    }

    public com.bandcamp.shared.network.GsonRequest a(Class p2, String p3, boolean p4)
    {
        com.bandcamp.shared.network.GsonRequest v2_2 = com.bandcamp.shared.network.API.a().a(p3, com.google.gson.reflect.TypeToken.get(p2));
        v2_2.b(p4);
        return v2_2;
    }

    public java.net.URL a(String p2)
    {
        return com.bandcamp.shared.network.API.a().f(p2);
    }

    public java.net.URL a(String p1, com.bandcamp.android.network.f p2)
    {
        java.util.List v2_1;
        if (p2 != null) {
            v2_1 = p2.a();
        } else {
            v2_1 = 0;
        }
        return this.a(p1, v2_1);
    }

    public void a(long p1)
    {
        com.bandcamp.shared.network.d.a(((int) p1));
        return;
    }

    public void a(java.util.Collection p4)
    {
        String v0_1 = com.bandcamp.fanapp.a.c().b();
        if (v0_1 != null) {
            p4.add(new android.util.Pair("admin_impersonate_fan", v0_1));
        }
        return;
    }

    public void a(java.util.Map p3)
    {
        String v0_1 = com.bandcamp.fanapp.a.c().b();
        if (v0_1 != null) {
            p3.put("admin_impersonate_fan", v0_1);
        }
        return;
    }

    public com.bandcamp.android.util.Promise b(long p2, char p4, long p5)
    {
        com.bandcamp.android.network.f v0_1 = new com.bandcamp.android.network.f();
        v0_1.a("band_id", Long.toString(p2));
        v0_1.a("tralbum_type", Character.toString(p4));
        v0_1.a("tralbum_id", Long.toString(p5));
        return this.a("/api/mobile/24/tralbum_details", 0, v0_1);
    }

    public com.bandcamp.android.util.Promise b(String p3)
    {
        java.util.HashMap v0_1 = new java.util.HashMap();
        v0_1.put("code", p3.trim());
        return this.a("/api/mobile/24/fan_yum", 1, v0_1);
    }

    public com.bandcamp.shared.network.e b(String p4, boolean p5, java.util.Collection p6)
    {
        com.bandcamp.shared.network.e v4_1 = com.bandcamp.shared.network.API.a().g(p4);
        v4_1.b(p5);
        String v5_2 = v4_1.e();
        String v6_2 = p6.iterator();
        while (v6_2.hasNext()) {
            String v0_3 = ((android.util.Pair) v6_2.next());
            boolean v1_0 = v0_3.second;
            if (!(v1_0 instanceof String)) {
                if (!(v1_0 instanceof Integer)) {
                    if (!(v1_0 instanceof Double)) {
                        if (!(v1_0 instanceof Float)) {
                            if (!(v1_0 instanceof Long)) {
                                if ((v1_0 instanceof Boolean)) {
                                    v5_2.put(((String) v0_3.first), ((Boolean) v1_0).booleanValue());
                                }
                            } else {
                                v5_2.put(((String) v0_3.first), ((Long) v1_0).longValue());
                            }
                        } else {
                            v5_2.put(((String) v0_3.first), ((double) ((Float) v1_0).floatValue()));
                        }
                    } else {
                        v5_2.put(((String) v0_3.first), ((Double) v1_0).doubleValue());
                    }
                } else {
                    v5_2.put(((String) v0_3.first), ((Integer) v1_0).intValue());
                }
            } else {
                v5_2.put(((String) v0_3.first), ((String) v1_0));
            }
        }
        v4_1.b("X-Requested-With", "com.bandcamp.android");
        return v4_1;
    }

    public com.bandcamp.shared.network.e b(String p2, boolean p3, org.json.JSONObject p4)
    {
        com.bandcamp.shared.network.e v2_1 = com.bandcamp.shared.network.API.a().g(p2);
        v2_1.b(p3);
        v2_1.a(p4);
        v2_1.b("X-Requested-With", "com.bandcamp.android");
        return v2_1;
    }

    public com.bandcamp.android.util.Promise c(String p3)
    {
        java.util.HashMap v0_1 = new java.util.HashMap(2);
        v0_1.put("address", p3);
        this.a(v0_1);
        return this.a("/api/mobile/24/linkpaypal", 1, v0_1);
    }
}
